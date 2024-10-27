#!/bin/sh

# important
sudo apt update
sudo apt install -y apt-transport-https ca-certificates software-properties-common gnupg lsb-release debian-keyring debian-archive-keyring curl

# go
sudo add-apt-repository -y ppa:longsleep/golang-backports
sudo apt update
sudo apt install -y golang-go

# nvim
sudo add-apt-repository ppa:neovim-ppa/stable
sudo apt update
sudo apt install -y neovim

# make
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test
sudo apt update
sudo apt install -y make

# fish
sudo apt-add-repository -y ppa:fish-shell/release-3
sudo apt update
sudo apt install -y fish

chsh -s /usr/bin/fish $USER

cat > ~/.config/fish/config.fish <<EOL
set fish_greeting
starship init fish | source
EOL

# starship
sudo apt install -y curl
curl -sS https://starship.rs/install.sh | sh -s -- -y

cat > ~/.config/starship.toml <<EOL
[line_break]
disabled = true
EOL

starship preset nerd-font-symbols >> ~/.config/starship.toml

# clone backend
cd
git clone https://github.com/zlatoivan/zlatoivan_ru.git

# caddy
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install -y caddy

rsync -a ~/zlatoivan_ru/Caddifile /etc/caddy/Caddyfile
sudo systemctl restart caddy

# docker
sudo curl -fsSL 'https://download.docker.com/linux/ubuntu/gpg' | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install -y docker-ce

# docker compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# ssh (убрать приветствие)
sudo chmod -x /etc/update-motd.d/*
