
export interface QuestionResponse {
    question: string;
    answers: Array<string>;
}

export interface AnswerResponse {
    isAnswerCorrect: boolean;
    hasNext: boolean;
    image: string;
}

class TriviaService {
    api: string

    constructor(api: string) {
        this.api = api;
    }

    getQuestion(): Promise<QuestionResponse> {
        return apiGet<QuestionResponse>(this.api + "/api/v1/trivia/question")
    }

    submitReponse(i: number): Promise<AnswerResponse> {
        const data = {
            index: i
        }
        return apiPost<AnswerResponse>(this.api + "/api/v1/trivia/question", data)
    }

    reset() {
        apiPost(this.api + "/api/v1/trivia/reset", null)
    }
}


function apiGet<T>(url: string): Promise<T> {
    return fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error(response.statusText)
            }
        return response.json() as Promise<T>
    }).then((responseT) => {return responseT});
}

function apiPost<T>(url: string, body: any): Promise<T> {
    return fetch(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: {
            "Content-Type": "application/json"
        } 
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(response.statusText)
            }
        return response.json() as Promise<T>
    }).then((responseT) => {return responseT});
}

export default TriviaService;
