# fiber-file-uploader

the simple file uploader

# demo

<img src="./img/image.gif" width="800px">

# quick start

- _$PWD/files_ set your directory

```bash
$ docker run -p 3000:3000 --rm -v $PWD:/app/files -it nguitar/fiber-file-uploader:1.00
```

# kubernetes sample

- Make the necessary pvc at the beginning.
- Create Service (NordPort or Ingress)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: fiber-file-uploader
  name: fiber-file-uploader
spec:
  replicas: 1
  selector:
    matchLabels:
      app: fiber-file-uploader
  template:
    metadata:
      labels:
        app: fiber-file-uploader
    spec:
      containers:
        - image: nguitar/fiber-file-uploader:1.00
          name: fiber-file-uploader
          volumeMounts:
            - mountPath: /app/files
              name: data
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: file-upload
```
