type generateUrlResponse  = {
    url: string
}

type generateUrlErrorResponse = {
    message: string
    error: string
}

class urlShotenerService {

    shortenerServiceUrl: string = "";
    constructor(shortenerServiceUrl: string) {
        this.shortenerServiceUrl = shortenerServiceUrl;
    }

    async generateUrl(url: string): Promise<string | null>{
        let generateShortUrlUri = `${this.shortenerServiceUrl}/api/createShortURL`;

        return await fetch(generateShortUrlUri, {
            method: 'POST',
            headers: {
                'content-type': 'application/json',
            },
            body: JSON.stringify(
                {
                    'url': url,
                }
            )
        }).then(async (response:Response) => {
            if(response.status == 200) {
                return response.json();
            }
            if(response.status == 400) {
                return response.json();
            }

            return null;
        }).then((data:generateUrlResponse | generateUrlErrorResponse) => {
            let urlResult:string = "";
            let error:string = "";
            if("url" in data) {
                urlResult = data.url;
            }
            if("error" in data && "message" in data) {
                error = data.error;
            }

            return [urlResult, error];
        }).catch((_) => {
            return null;
        })
    }
}


export default urlShotenerService;