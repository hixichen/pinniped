// Copyright 2020 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientauthenticationv1beta1 "k8s.io/client-go/pkg/apis/clientauthentication/v1beta1"

	"github.com/suzerain-io/pinniped/generated/1.19/apis/pinniped/v1alpha1"
	"github.com/suzerain-io/pinniped/internal/testutil"
)

func TestExchangeToken(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("invalid configuration", func(t *testing.T) {
		t.Parallel()
		got, err := ExchangeToken(ctx, "", "", "")
		require.EqualError(t, err, "could not get API client: invalid configuration: no configuration has been provided, try setting KUBERNETES_MASTER environment variable")
		require.Nil(t, got)
	})

	t.Run("server error", func(t *testing.T) {
		t.Parallel()
		// Start a test server that returns only 500 errors.
		caBundle, endpoint := testutil.TLSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("some server error"))
		})

		got, err := ExchangeToken(ctx, "", caBundle, endpoint)
		require.EqualError(t, err, `could not login: an error on the server ("some server error") has prevented the request from succeeding (post credentialrequests.pinniped.dev)`)
		require.Nil(t, got)
	})

	t.Run("login failure", func(t *testing.T) {
		t.Parallel()
		// Start a test server that returns success but with an error message
		errorMessage := "some login failure"
		caBundle, endpoint := testutil.TLSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			_ = json.NewEncoder(w).Encode(&v1alpha1.CredentialRequest{
				TypeMeta: metav1.TypeMeta{APIVersion: "pinniped.dev/v1alpha1", Kind: "CredentialRequest"},
				Status:   v1alpha1.CredentialRequestStatus{Message: &errorMessage},
			})
		})

		got, err := ExchangeToken(ctx, "", caBundle, endpoint)
		require.EqualError(t, err, `login failed: some login failure`)
		require.Nil(t, got)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		expires := metav1.NewTime(time.Now().Truncate(time.Second))

		// Start a test server that returns successfully and asserts various properties of the request.
		caBundle, endpoint := testutil.TLSTestServer(t, func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodPost, r.Method)
			require.Equal(t, "/apis/pinniped.dev/v1alpha1/credentialrequests", r.URL.Path)
			require.Equal(t, "application/json", r.Header.Get("content-type"))

			body, err := ioutil.ReadAll(r.Body)
			require.NoError(t, err)
			require.JSONEq(t,
				`{
				  "kind": "CredentialRequest",
				  "apiVersion": "pinniped.dev/v1alpha1",
				  "metadata": {
					"creationTimestamp": null
				  },
				  "spec": {
					"type": "token",
					"token": {}
				  },
				  "status": {}
				}`,
				string(body),
			)

			w.Header().Set("content-type", "application/json")
			_ = json.NewEncoder(w).Encode(&v1alpha1.CredentialRequest{
				TypeMeta: metav1.TypeMeta{APIVersion: "pinniped.dev/v1alpha1", Kind: "CredentialRequest"},
				Status: v1alpha1.CredentialRequestStatus{
					Credential: &v1alpha1.CredentialRequestCredential{
						ExpirationTimestamp:   expires,
						ClientCertificateData: "test-certificate",
						ClientKeyData:         "test-key",
					},
				},
			})
		})

		got, err := ExchangeToken(ctx, "", caBundle, endpoint)
		require.NoError(t, err)
		require.Equal(t, &clientauthenticationv1beta1.ExecCredential{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ExecCredential",
				APIVersion: "client.authentication.k8s.io/v1beta1",
			},
			Status: &clientauthenticationv1beta1.ExecCredentialStatus{
				ClientCertificateData: "test-certificate",
				ClientKeyData:         "test-key",
				ExpirationTimestamp:   &expires,
			},
		}, got)
	})
}