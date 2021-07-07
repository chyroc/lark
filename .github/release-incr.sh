set -ex

cur_tag=`git describe --tags --abbrev=0`

last_ver=`echo $cur_tag | cut -d '.' -f 3-3`

new_ver=$((last_ver+1))

new_tag="v0.0.$new_ver"

cat <<EOT > version.go
package lark

const version = "$new_tag"
EOT

git commit -a -m "release: $new_tag"

git tag $new_tag

echo "$cur_tag, $last_ver, $new_ver, $new_tag"

echo "success"
