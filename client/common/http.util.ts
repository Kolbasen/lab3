import Axios, { AxiosInstance } from 'axios';

export class HttpCllient {
    private api: AxiosInstance;
    constructor(url: string) {
        this.api = Axios.create({
            baseURL: url,
        })
    }

    public async get<T>(endpoint: string) {
        const response = await this.api.get<T>(endpoint);
        return response.data;
    }

    public async post<T, K>(endpoint: string, data: K) {
        const response = await this.api.post<T>(endpoint, data);
        return response.data;
    }
}

