import EmailPassword from "supertokens-auth-react/recipe/emailpassword";
import Session from "supertokens-auth-react/recipe/session";

export function getAppName() {
    return process.env.SUPERTOKENS_APP_NAME || '';
}

export function getApiDomain() {
    return process.env.SUPERTOKENS_API_URL || '';
}

export function getWebsiteDomain() {
    return process.env.SUPERTOKENS_WEBSITE_URL || '';
}

export function getApiBasePath() {
    return process.env.SUPERTOKENS_API_BASE_PATH || '';
}

export function getApiGatewayPath() {
    return process.env.SUPERTOKENS_API_GATEWAY_PATH || '';
}

export const SuperTokensConfig = {
    appInfo: {
        appName: getAppName(),
        apiDomain: getApiDomain(),
        websiteDomain: getWebsiteDomain(),
        apiBasePath: getApiBasePath(),
        apiGatewayPath: getApiGatewayPath(),
    },
    // recipeList contains all the modules that you want to
    // use from SuperTokens. See the full list here: https://supertokens.com/docs/guides
    recipeList: [EmailPassword.init(), Session.init()],
    enableDebugLogs: true,
};

export const recipeDetails = {
    docsLink: "https://supertokens.com/docs/emailpassword/introduction",
};
