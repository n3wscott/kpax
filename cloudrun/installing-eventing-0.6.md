This is a simplified verison of https://github.com/knative/docs/blob/master/docs/install/Knative-with-GKE.md

## Before you Begin:

1. [Installing the Google Cloud SDK and `kubectl`](https://github.com/knative/docs/blob/master/docs/install/Knative-with-GKE.md#installing-the-google-cloud-sdk-and-kubectl)
1. [Setting environment variables](https://github.com/knative/docs/blob/master/docs/install/Knative-with-GKE.md#setting-environment-variables)
1. [Setting up a Google Cloud Platform project](https://github.com/knative/docs/blob/master/docs/install/Knative-with-GKE.md#setting-up-a-google-cloud-platform-project)

_Skip_ [Creating a Kubernetes cluster](https://github.com/knative/docs/blob/master/docs/install/Knative-with-GKE.md#creating-a-kubernetes-cluster)

## Installing Knative Eventing with Cloud Run on GKE

The following commands install all available Knative Eventing components:

1. To install Knative, first install the CRDs by running the `kubectl apply`
   command once with the `-l knative.dev/crd-install=true` flag. This prevents
   race conditions during the install, which cause intermittent errors:

   ```bash
   kubectl apply --selector knative.dev/crd-install=true \
   --filename https://github.com/knative/eventing/releases/download/v0.6.0/release.yaml \
   --filename https://github.com/knative/eventing-sources/releases/download/v0.6.0/eventing-sources.yaml
   ```

1. To complete the install of Knative and its dependencies, run the
   `kubectl apply` command again, this time without the `--selector` flag, to
   complete the install of Knative and its dependencies:

   ```bash
   kubectl apply \
   --filename https://github.com/knative/eventing/releases/download/v0.6.0/release.yaml \
   --filename https://github.com/knative/eventing-sources/releases/download/v0.6.0/eventing-sources.yaml
   ```

1. Monitor the Knative components until all of the components show a `STATUS` of
   `Running`:

   ```bash
   kubectl get pods --namespace knative-eventing
   kubectl get pods --namespace knative-sources
   ```

## What's next

Now that your cluster has Knative Eventing installed, you can see what Knative
Eventing has to offer.

To get started with Knative Eventing, pick one of the
[Eventing Samples](https://github.com/knative/docs/tree/master/docs/eventing/samples)
to walk through.
