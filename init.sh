#!/bin/sh

# add-apt-repository
sudo apt-get install software-properties-common

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

# caddy
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy

rsync -a ~/zlatoivan/Caddifile /etc/caddy/Caddyfile
sudo systemctl restart caddy

# clone backend
cd
git clone https://github.com/zlatoivan/zlatoivan_ru.git

# ssh (убрать приветствие)
sudo chmod -x /etc/update-motd.d/*
