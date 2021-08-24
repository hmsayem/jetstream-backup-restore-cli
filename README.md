## Disaster Recovery CLI for the JetStream System

##### Install Cert Manager
```
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.5.0/cert-manager.yaml
```
##### Create certificates
```
kubectl apply -f certs/nats-certs.yml
```
####  Run Account Server
##### Create Configmaps
```
kubectl apply -f account-server/configmaps/sys-op.yml
```
##### Create Secrets
```
kubectl apply -f account-server/secrets/sys.yml
```
##### Install Account Server with Helm
```
helm install stan-account-server account-server/helm/account-server    
```
##### Port Forwarding
```
kubectl port-forward svc/stan-account-server  9090
```
##### Push Accounts
```
go run account-server/accounts/push.go
```
####  Run Jetstream System 
##### Install NATS with Helm
```
helm install stan jetstream/helm/nats    
```
##### Port Forwarding
```
kubectl port-forward svc/stan-nats  4222
```

#### Backup
##### With Client Certificate Authentication
```
go run main.go backup \
--target <stream> \
--dir <backup_directory> \
--tlscert <client_certificate_path> \
--tlskey <tls_key_path> \
--tlsca <ca_certificate_path> \
--user <username> \
--password <password> \
--creds <user_credential>
```
##### Without Client Certificate Authentication
```
go run main.go backup \
--target <stream> \
--dir <backup_directory> \
--tls=false \
--user <username> \
--password <password> \
--creds <user_credential>
```
#### Restore
##### With Client Certificate Authentication
```
go run main.go restore \
--target <stream> \
--dir <backup_directory> \
--tlscert <client_certificate_path> \
--tlskey <tls_key_path> \
--tlsca <ca_certificate_path> \
--user <username> \
--password <password> \
--creds <user_credential>
```
##### Without Client Certificate Authentication
```
go run main.go restore \
--target <stream> \
--dir <backup_directory> \
--tls=false \
--user <username> \
--password <password> \
--creds <user_credential>
```
