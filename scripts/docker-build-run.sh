docker run -p 3000:3000 --name metapier-server --rm -it $(docker build -q -t metapier/server -f build/package/Dockerfile .)