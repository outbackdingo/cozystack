export NAME=nfs-driver
export NAMESPACE=cozy-$(NAME)

include ../../../scripts/common-envs.mk
include ../../../scripts/package.mk

update:
	rm -rf charts
	helm repo add csi-driver-nfs https://raw.githubusercontent.com/kubernetes-csi/csi-driver-nfs/master/charts
	helm repo update csi-driver-nfs
	helm pull csi-driver-nfs/csi-driver-nfs --untar --untardir charts
