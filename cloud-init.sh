#!/bin/sh

# go +
sudo snap install go --classic

# nvim +
sudo snap install nvim --classic

# make +
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test
sudo apt update
sudo apt install make

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

if [ -f ~/.config/starship.toml ]; then
    starship preset nerd-font-symbols >> ~/.config/starship.toml
else
    echo "⛔⛔⛔ Файл ~/.config/starship.toml не существует. Команда не выполнена."
fi

# caddy
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy

domain="zlatoivan.ru"
proxy="localhost:7070"
cat > /etc/caddy/Caddyfile <<EOL
$domain {
    reverse_proxy $proxy {
        header_up X-Real-IP {}
        header_up X-Forwarded-For {http.request.header.X-Forwarded-For}
        header_up X-Forwarded-Port {http.request.header.X-Forwarded-Port}
    }
    log {
        output file /var/log/caddy/access.log {
            roll_size 10MB
            roll_keep 100
            roll_keep_days 30
        }
        level INFO
    }
}
EOL
sudo systemctl restart caddy

# clone backend
cd
git clone https://github.com/zlatoivan/zlatoivan_ru.git

# ssh
cat >> /etc/ssh/sshd_config <<EOL
Banner none               # Приветствие
PrintLastLog no           # Информация о последнем сеансе
ClientAliveInterval 1800  # 30 минут в секундах
ClientAliveCountMax 4     # Максимум 4 попытки
EOL

sudo systemctl restart ssh || echo "⛔⛔⛔Failed to restart SSH"