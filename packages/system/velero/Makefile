export NAME=vertical-pod-autoscaler
export NAMESPACE=cozy-$(NAME)

include ../../../scripts/package.mk

update:
	rm -rf charts
	# VirtualPodAutoscaler operator
	helm repo add tanzu https://vmware-tanzu.github.io/helm-charts
	helm repo update tanzu
	helm pull tanzu/velero --untar --untardir charts
