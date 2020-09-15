# kubebuilder-test
Kubebuilder Test

## Init kubebuilder test
```
1. Initialize a new project
kubebuilder init --domain zmz.example.org --owner "zmz" --repo github.com/Zheng-Mz/kubebuilder-test

2.kubebuilder create api will prompt the user asking if it should scaffold the Resource and / or Controller.
kubebuilder create api --group zmzapp --version v1 --kind TestK
```