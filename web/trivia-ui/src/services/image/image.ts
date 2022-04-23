
export interface QuestionResponse {
    question: string;
    answers: Array<string>;
}

export interface AnswerResponse {
    isAnswerCorrect: boolean;
    hasNext: boolean;
}

class ImageService {
    api: string

    constructor(api: string) {
        this.api = api;
    }

    getImage(imagename: string): Promise<Blob> {
        var url = new URL(this.api + "/api/v1/photo")
        var params: string[][] = [
            ["name", imagename],
            ["width", window.innerWidth.toString()],
            ["height", window.innerHeight.toString()]
        ]
        url.search = new URLSearchParams(params).toString();

        return apiGet(url.toString())
    }
}


function apiGet(url: string): Promise<Blob> {
    console.log(url)
    return fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(response.statusText)
            }
        return response.blob() 
    });
}

export default ImageService;
