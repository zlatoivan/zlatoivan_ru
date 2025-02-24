set fish_greeting
starship init fish | source

set -gx GOPATH $HOME/go
# set -gx PATH $PATH $GOPATH/bin

# set -gx GOPATH /root/go
set -gx PATH /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin

# Aliases

alias drma  'docker rm -v -f $(docker ps -qa)'
alias drme  'docker rm -v $(docker ps --filter status=exited -q)'

alias gco   'git checkout'
alias gcb   'git checkout -b'
alias gcm   'git checkout master'
alias gb    'git branch'
alias gbd   'git branch -d'
alias gbD   'git branch -D'

alias gs    'git status'
alias ga    'git add'
alias gc    'git commit -m'
alias gac   'git add . && git commit -m'
alias gaca  'git add . && git commit --amend --no-edit'

alias gp    'git push'
alias gpf   'git push -f'
alias gpu   'git push -u origin'
alias gpd   'git push -d origin'

alias gf    'git fetch'
alias grb   'git rebase'
alias grbm  'git rebase origin/master'
alias grbim 'git rebase -i origin/master'
alias grbc  'git rebase --continue'

alias gr    'git reset'
alias grs   'git reset --soft'
alias grh   'git reset --hard'
alias gr1   'git reset HEAD~1'
alias gcl   'git clean -fd'

alias gl    'git log'
alias glp   'git log --pretty=format:"%C(yellow)%h %C(red)%ad %C(blue)%an %C(default)%s%C(green)%d" --date=format:"%d.%m.%y %H:%M"'
alias glpa  'git log --pretty=format:"%C(yellow)%h %C(red)%ad %C(blue)%an %C(default)%s%C(green)%d" --date=format:"%d.%m.%y %H:%M" --all'
alias glpg  'git log --pretty=format:"%C(yellow)%h %C(red)%ad %C(blue)%an %C(default)%s%C(green)%d" --date=format:"%d.%m.%y %H:%M" --graph'
alias glpd  'git log --pretty=format:"%C(yellow)%h %C(red)%ad %C(blue)%an %C(default)%s%C(green)%d" --date=format:"%d.%m.%y %H:%M" --graph --left-right --cherry-pick HEAD...@{u}'
alias gld   'git log --oneline --graph --left-right --cherry-pick HEAD...@{u}'
alias glo   'git log --oneline'

alias gpl   'git pull'
alias gcp   'git cherry-pick'
alias gcv   'git cherry -v --abbrev=8 origin/master'
alias gbv   'git branch -vv'
alias gds   'git diff --stat'
alias gdss  'git diff --shortstat'