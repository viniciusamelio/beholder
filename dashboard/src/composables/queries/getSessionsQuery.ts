import { env } from "@/env";
import type { Session } from "@/types";
import { useQuery } from "@pinia/colada"


export const getSessionsQuery = (id: number) => { 
    return useQuery(
        {
            key: () => ["get-sessions", id],
            query: async () => {
                const response = await fetch(`${env.baseUrl}/environment/${id}/sessions`);
                return await response.json() as Session[];
            }
        }
    )
}