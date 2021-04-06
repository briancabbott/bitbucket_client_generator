cd bitbucket_client

NEW_VERSION=v0.1.1 

git add . 
git commit -m "publishing updated BitBucket client: $NEW_VERSION"
git push
git tag $NEW_VERSION
git push origin $NEW_VERSION

NEW_VERSION=v0.1.1; GOPROXY=proxy.golang.org go list -m github.com/briancabbott/bitbucket_client@$NEW_VERSION
cd .. 

