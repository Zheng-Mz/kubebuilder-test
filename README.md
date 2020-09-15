# kubebuilder-test
Kubebuilder Test

## Init kubebuilder test
```
1. Initialize a new project
kubebuilder init --domain zmz.example.org --owner "zmz" --repo github.com/Zheng-Mz/kubebuilder-test

2. Kubebuilder create api will prompt the user asking if it should scaffold the Resource and / or Controller.
kubebuilder create api --group zmzapp --version v1 --kind TestK

3. Install CRD
make install

4. Run controller
make run

5. Uninstall CRD
make uninstall

6. Build Contrller image
make docker-build [IMG=controller-image-name:tag]

7. Deploy controller in the configured Kubernetes cluster in ~/.kube/config.(Create crd/deploy controller)
make deploy
```