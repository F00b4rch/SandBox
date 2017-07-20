take info from openrc.sh and add it to env in dockerfile
then :
`sudo docker build . --tag="nova:latest"`

in your zshrc/bashrc :
`alias nova='sudo docker run -it nova'`

then :
`nova`
