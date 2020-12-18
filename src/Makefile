deployCF:
	gcloud functions deploy hello --entry-point HelloWorld --runtime go113 --trigger-http --region us-east1 --set-env-vars MSG="Cloud Functions" --allow-unauthenticated --source .
build:
	pack build --builder gcr.io/buildpacks/builder:v1 --env GOOGLE_FUNCTION_SIGNATURE_TYPE=http --env GOOGLE_FUNCTION_TARGET=HelloWorld  hello-fn
run:
	docker run --rm -p 8080:8080 hello-fn
buildGCR:
	gcloud builds submit --tag gcr.io/openstep-website/hello-fn
deployCloudRun:
	gcloud run deploy --image gcr.io/openstep-website/hello-fn --set-env-vars MSG="Cloud Run" --platform managed
