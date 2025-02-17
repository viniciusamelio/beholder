import { env } from '@/env';
import type { Environment } from '@/types';
import { useQuery } from '@pinia/colada';

export const getEnvironmentsQuery = () => { 
    return useQuery({
        key: () => ["get-environments"],
        query: async () => {
            const response = await fetch(`${env.baseUrl}/environment`);
            return await response.json() as Environment[];
        }
    });
}