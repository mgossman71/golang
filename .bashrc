#export KUBECONFIG=~/.kube/pi01
unset KUBECONFIG

HISTSIZE=200000
HISTFILESIZE=$HISTSIZE
HISTCONTROL=ignoreboth:erasedups
shopt -s histappend

_bash_history_sync() {
  builtin history -a         #1
  HISTFILESIZE=$HISTSIZE     #2
  builtin history -c         #3
  builtin history -r         #4
}

history() {                  #5
  _bash_history_sync
  builtin history "$@"
}

# Git - prompt shows active branch.
export PS1="\h:\\w:\$(git branch 2>/dev/null | grep '^*' | colrm 1 2)\$ "

# cluster-select
alias cluster-select='cd ~/.kube; . cluster-sel.sh; cd - 1>/dev/null'

# pod-stats
alias pod-stats='kubectl get pods --all-namespaces -o wide'

# watch
watch() {
while true
do
  clear
  date
  $1
  sleep $2
done
}

scale-apps() {
kubectl scale -n $1 --replicas=$2 deployment/sonarr-deployment
kubectl scale -n $1 --replicas=$2 deployment/couchpotato-deployment
kubectl scale -n $1 --replicas=$2 deployment/sabnzbd-deployment
kubectl scale -n $1 --replicas=$2 deployment/gozznet-old 
}
