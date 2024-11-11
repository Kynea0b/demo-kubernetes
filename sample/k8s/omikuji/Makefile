export SOPS_AGE_KEY_FILE=$(HOME)/work/microservice-demo/k8s/omikuji/keys.txt

apply:
	helm secrets upgrade omikuji . --install \
	-f values.yaml \
	-f secrets/secrets.enc.yaml

delete:
	helm uninstall omikuji
	kubectl delete pvc data-omikuji-mysql-0

dbpass:
	kubectl get secret --namespace default omikuji-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo

encrypt:
	sops --encrypt --age age158hhkp5kn6aeauhqkaqf2rqcqvjxan0mkuz30qkcr4g3sf9skpzqkm07ad secrets/secrets.yaml > secrets/secrets.enc.yaml
