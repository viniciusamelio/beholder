import { env } from "@/env";
import { SessionRequestsSchema } from "@/schema/schema_pb";
import { fromBinary } from "@bufbuild/protobuf";
import { useQuery } from "@pinia/colada";

export const useGetSessionQuery = (id: string) => {
    const query = async (id: string) => {
        const response = await fetch(`${env.baseUrl}/sessions/${id}`);
        const bytes = await response.bytes();
        const value = fromBinary(SessionRequestsSchema, bytes);
        return value;
    }
    
    return useQuery({
        key: () => ["get-session", id],
        query: () => query(id)
    });
}