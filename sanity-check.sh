# Load environment variables.
source .env 

# Test connectivity to the eMASS API.
#curl \
#    -L "$EMASS_API_URL/api" \
#    --cert "$NPE_CERTIFICATE_FILEPATH" \
#    --key "$NPE_KEY_FILEPATH" &&\
#echo 

# Get information about a specific system.
#curl \
#    -L "$EMASS_API_URL/api/systems/$EMASS_SYSTEM_ID" \
#    --cert "$NPE_CERTIFICATE_FILEPATH" \
#    --key "$NPE_KEY_FILEPATH" \
#    -H "api-key: $EMASS_API_KEY" &&\
#echo 

# Get devices within a specific system.
curl \
    -L "$EMASS_API_URL/api/systems/$EMASS_SYSTEM_ID/devices" \
    --cert "$NPE_CERTIFICATE_FILEPATH" \
    --key "$NPE_KEY_FILEPATH" \
    -H "api-key: $EMASS_API_KEY" &&\
echo 
