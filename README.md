## Disaster Recovery CLI for the JetStream System Running on Kubernetes Cluster

### Prerequisite
##### Install Cert Manager
```
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.5.0/cert-manager.yaml
```
##### Create certificates
```
kubectl apply -f certs/nats-certs.yml
```

###  Deploy NATS Cluster with Different Authentications
##### Basic Authentication
```bash
# Add nats chart registry
$ helm repo add nats https://nats-io.github.io/k8s/helm/charts/ 
# Update helm registries
$ helm repo update
# Install nats/nats chart into demo namespace
$ helm install nats nats/nats -n demo \
--set nats.jetstream.enabled=true \
--set nats.jetstream.fileStorage.enabled=true \
--set cluster.enabled=true \
--set cluster.replicas=3 \
--set auth.enabled=true \
--set-string auth.basic.users[0].user="sample-user",auth.basic.users[0].password="changeit"
```
##### JWT Authentication
```bash
# Add nats chart registry
$ helm repo add nats https://nats-io.github.io/k8s/helm/charts/ 
# Update helm registries
$ helm repo update
# Install nats/nats chart into demo namespace
$ helm install nats nats/nats -n demo \
--set nats.jetstream.enabled=true \
--set nats.jetstream.fileStorage.enabled=true \
--set cluster.enabled=true \
--set cluster.replicas=3 \
--set auth.enabled=true \
--set auth.resolver.type=full \
--set auth.resolver.operator=eyJ0eXAiOiJKV1QiLCJhbGciOiJlZDI1NTE5LW5rZXkifQ.eyJhdWQiOiJPQU5US0NDTkFQTUdCNE9YRE1YT1ZON01DQUVKNVZGV1ZXVEFVVlVXQllBUFhMWlpHU1NZVTRLVCIsImV4cCI6MTk0NTQyNjA0MSwianRpIjoiWktRTllaTlNRTUhNQzNHREdVRVpDUFlDT0RNSjIyMzRPM0pGTjUzWlZYWEZBUFU3Qlg2QSIsImlhdCI6MTYyOTg5MzI0MSwiaXNzIjoiT0FOVEtDQ05BUE1HQjRPWERNWE9WTjdNQ0FFSjVWRldWV1RBVVZVV0JZQVBYTFpaR1NTWVU0S1QiLCJuYW1lIjoiS08iLCJuYmYiOjE2Mjk4OTMyNDEsInN1YiI6Ik9BTlRLQ0NOQVBNR0I0T1hETVhPVk43TUNBRUo1VkZXVldUQVVWVVdCWUFQWExaWkdTU1lVNEtUIiwibmF0cyI6eyJzaWduaW5nX2tleXMiOlsiT0FOVEtDQ05BUE1HQjRPWERNWE9WTjdNQ0FFSjVWRldWV1RBVVZVV0JZQVBYTFpaR1NTWVU0S1QiXSwidHlwZSI6Im9wZXJhdG9yIiwidmVyc2lvbiI6Mn19.jxs4znpE50PzRFfKOjENlFQTfsRHH5VqIplnTgAziUJuYBSNmBQeYsBDJTOgLJyADgtqIWkAQF_G5K7xuVXpCg \
--set auth.resolver.systemAccount=ABQBM7PTUNWRWQRWFFQGRCVRQ7ULYSZQLCMGDK62WYRHRO3NUN3SUONF \
--set auth.resolver.resolverPreload.ABQBM7PTUNWRWQRWFFQGRCVRQ7ULYSZQLCMGDK62WYRHRO3NUN3SUONF=eyJ0eXAiOiJKV1QiLCJhbGciOiJlZDI1NTE5LW5rZXkifQ.eyJqdGkiOiJDRTZHVUdISEYzVzRGUU5VMjdVWks2SUtNSkpQWEEzUlEyNDVYTE5LWEVIS0xaQ1ZaT1FBIiwiaWF0IjoxNjI5ODkzMjQxLCJpc3MiOiJPQU5US0NDTkFQTUdCNE9YRE1YT1ZON01DQUVKNVZGV1ZXVEFVVlVXQllBUFhMWlpHU1NZVTRLVCIsIm5hbWUiOiJTWVMiLCJzdWIiOiJBQlFCTTdQVFVOV1JXUVJXRkZRR1JDVlJRN1VMWVNaUUxDTUdESzYyV1lSSFJPM05VTjNTVU9ORiIsIm5hdHMiOnsibGltaXRzIjp7InN1YnMiOi0xLCJkYXRhIjotMSwicGF5bG9hZCI6LTEsImltcG9ydHMiOi0xLCJleHBvcnRzIjotMSwid2lsZGNhcmRzIjp0cnVlLCJjb25uIjotMSwibGVhZiI6LTF9LCJkZWZhdWx0X3Blcm1pc3Npb25zIjp7InB1YiI6e30sInN1YiI6e319LCJ0eXBlIjoiYWNjb3VudCIsInZlcnNpb24iOjJ9fQ.DM8U4Ld4OWmd-hk9fobMI3fWMDZdLr-Q33Uq7h5eoM-RMVKN-5nuUlmaxPffYRwE1egVIn9mQuu7YYmwX31wDQ \
--set auth.resolver.resolverPreload.ADFC3YLAU56N26HGN7TGPWDXQNZSBQEZKQXPPRP24HK46VJYQX45S2UH=eyJ0eXAiOiJKV1QiLCJhbGciOiJlZDI1NTE5LW5rZXkifQ.eyJqdGkiOiJEMlYzVE9NUTQ1WjRGUFNTM0hCMk5TR1dLV0xITDVNSUpVWENNSTRGVEtYNTIzNjdWRUhRIiwiaWF0IjoxNjI5ODkzMjQxLCJpc3MiOiJPQU5US0NDTkFQTUdCNE9YRE1YT1ZON01DQUVKNVZGV1ZXVEFVVlVXQllBUFhMWlpHU1NZVTRLVCIsIm5hbWUiOiJYIiwic3ViIjoiQURGQzNZTEFVNTZOMjZIR043VEdQV0RYUU5aU0JRRVpLUVhQUFJQMjRISzQ2VkpZUVg0NVMyVUgiLCJuYXRzIjp7ImV4cG9ydHMiOlt7Im5hbWUiOiJ4LkV2ZW50cyIsInN1YmplY3QiOiJ4LkV2ZW50cyIsInR5cGUiOiJzdHJlYW0ifSx7Im5hbWUiOiJ4Lk5vdGlmaWNhdGlvbnMiLCJzdWJqZWN0IjoieC5Ob3RpZmljYXRpb25zIiwidHlwZSI6InNlcnZpY2UiLCJyZXNwb25zZV90eXBlIjoiU3RyZWFtIn1dLCJsaW1pdHMiOnsic3VicyI6LTEsImRhdGEiOi0xLCJwYXlsb2FkIjotMSwiaW1wb3J0cyI6LTEsImV4cG9ydHMiOi0xLCJ3aWxkY2FyZHMiOnRydWUsImNvbm4iOi0xLCJsZWFmIjotMSwibWVtX3N0b3JhZ2UiOi0xLCJkaXNrX3N0b3JhZ2UiOi0xLCJzdHJlYW1zIjotMSwiY29uc3VtZXIiOi0xfSwiZGVmYXVsdF9wZXJtaXNzaW9ucyI6eyJwdWIiOnt9LCJzdWIiOnt9fSwidHlwZSI6ImFjY291bnQiLCJ2ZXJzaW9uIjoyfX0.oXatnt7Tqt1iHDpUAKGroac9Sv6G4kbAPIt75BrBRh6B9MOFa_y8QLsUnIffI4-aG31cVYjECs7QlsNTPJ-oCg
```
##### Nkey Authentication
```bash
# Add nats chart registry
$ helm repo add nats https://nats-io.github.io/k8s/helm/charts/ 
# Update helm registries
$ helm repo update
# Install nats/nats chart into demo namespace
$ helm install nats nats/nats -n demo \
--set nats.jetstream.enabled=true \
--set nats.jetstream.fileStorage.enabled=true \
--set cluster.enabled=true \
--set cluster.replicas=3 \
--set auth.enabled=true \
--set auth.nkeys.users[0].nkey="UAWGGVEHXIEOWYVBN7RO7IIKIXHIOK6JWWUDJOWIUZ6L3XMIWM5IFGPD"
```
##### Token Authentication
```bash
# Add nats chart registry
$ helm repo add nats https://nats-io.github.io/k8s/helm/charts/ 
# Update helm registries
$ helm repo update
# Install nats/nats chart into demo namespace
$ helm install nats nats/nats -n demo \
--set nats.jetstream.enabled=true \
--set nats.jetstream.fileStorage.enabled=true \
--set cluster.enabled=true \
--set cluster.replicas=3 \
--set auth.enabled=true \
--set auth.token="secret"
```
##### TLS Authentication
```bash
# Add nats chart registry
$ helm repo add nats https://nats-io.github.io/k8s/helm/charts/ 
# Update helm registries
$ helm repo update
# Install nats/nats chart into demo namespace
$ helm install nats nats/nats -n demo \
--set nats.jetstream.enabled=true \
--set nats.jetstream.fileStorage.enabled=true \
--set nats.tls.secret.name=nats-client-tls \
--set nats.tls.ca="ca.crt" \
--set nats.tls.cert="tls.crt" \
--set nats.tls.key="tls.key" \
--set nats.tls.verify=true \
--set cluster.tls.secret.name=nats-server-tls \
--set cluster.tls.ca="ca.crt" \
--set cluster.tls.cert="tls.crt" \
--set cluster.tls.key="tls.key" \
--set cluster.enabled=true \
--set cluster.replicas=3 
```

### Port Forwarding
`kubectl port-forward` allows you to access and interact with internal Kubernetes cluster processes from your localhost.
```
kubectl port-forward svc/nats 4222
```

### Backup
```
go run main.go backup \
--target <stream> \
--dir <backup_directory> \
--tls <bool> \
--tlscert <client_certificate_path> \
--tlskey <tls_key_path> \
--tlsca <ca_certificate_path> \
--user <username> \
--password <password> \
--creds <user_credential>
```
### Restore
```
go run main.go restore \
--target <stream> \
--dir <backup_directory> \
--tls <bool> \
--tlscert <client_certificate_path> \
--tlskey <tls_key_path> \
--tlsca <ca_certificate_path> \
--user <username> \
--password <password> \
--creds <user_credential>
```
