package network

import (
  "os"
  "path/filepath"

  "github.com/openshift/cluster-network-operator/pkg/render"
  "github.com/pkg/errors"
  uns "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// renderDHCPDaemon returns the manifests of the DHCP Daemon
func renderDHCPDaemon(manifestDir string) ([]*uns.Unstructured, error) {
  objs := []*uns.Unstructured{}

  // render the manifests on disk
  data := render.MakeRenderData()
  data.Data["ReleaseVersion"] = os.Getenv("RELEASE_VERSION")
  data.Data["CNIPluginsSupportedImage"] = os.Getenv("CNI_PLUGINS_SUPPORTED_IMAGE")

  manifests, err := render.RenderDir(filepath.Join(manifestDir, "network/dhcp"), &data)
  if err != nil {
    return nil, errors.Wrap(err, "failed to render dhcp manifests")
  }
  objs = append(objs, manifests...)
  return objs, nil
}
