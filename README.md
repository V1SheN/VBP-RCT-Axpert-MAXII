# VBP-RCT-Axpert-MAXII


# Commands

```
docker-compose build voltronic-mqtt
```
```
docker-compose up -d --force-recreate voltronic-mqtt
```
# Axpert CLI 
```
rsync -avz /home/fish/Software/Development/github/VBP-RCT-Axpert-MAXII/go_inverter_cli/ fish@192.168.31.218:/opt/go_inverter_cli/

```

upload repo
```
rsync -avz /home/fish/Software/Development/github/Home-Assistant/docker-voltronic-homeassistant-master/ fish@192.168.31.218:/opt/docker-voltronic-homeassistant-master/
```

## Run CLI
```
docker exec -it voltronic-mqtt /opt/inverter-cli/bin/inverter_poller -d -1
```


```
https://raw.githubusercontent.com/svalouch/python-rctclient/master/rctclient/client.py
```
