set -e

keyvault=$1
keyversion=$2
signEndpoint="https://${keyvault}.vault.azure.net/keys/signing-key/${keyversion}/sign?api-version=7.1"
verifyEndpoint="https://${keyvault}.vault.azure.net/keys/signing-key/${keyversion}/verify?api-version=7.1"

verificationMessage="hello world"

echo "Verification Message: " $verificationMessage

digest=$(echo -n "${verificationMessage}" | sha256sum | cut -d ' ' -f 1 | xxd -r -p | base64)
echo "Digest: " $digest

signret=$(az rest \
	--url ${signEndpoint} \
	--resource https://vault.azure.net \
	--method post \
	--headers "Content-Type=application/json" \
	--body "{\"value\": \"${digest}\", \"alg\": \"ES256\"}")

signValue=$(echo $signret | jq -r .value)

echo "Signature IEEE 1363 " $signValue

verifyResult=$(az rest \
	--url ${verifyEndpoint} \
	--resource https://vault.azure.net \
	--method post \
	--headers "Content-Type=application/json" \
	--body "{\"digest\":\"${digest}\",\"value\": \"${signValue}\", \"alg\": \"ES256\"}")

echo "Verify Result: " $verifyResult

asn1=$(go run verify/verify.go --ieee1363 ${signValue})

echo "Signatire ASN.1 " $asn1