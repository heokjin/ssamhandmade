"# ssamhandmade" 

docker build -t heokjin/ssamhandmade .  <br>
docker run -p 80:80 -it --rm --name ssamhandmade ssamhandmade:latest
<p>
docker login <br>
docker push heokjin/ssamhandmade:latest