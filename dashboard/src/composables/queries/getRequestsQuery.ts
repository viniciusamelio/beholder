import { env } from "@/env";
import { useQuery } from "@pinia/colada"


export const useGetRequestsQuery = (id: number) => { 
    return useQuery({
        key: () => ["get-requests", id],
        query: async () => {
            const response = await fetch(`${env.baseUrl}/environment/${id}/requests`);
            return await response.json() as Request[];
        }
    })
}