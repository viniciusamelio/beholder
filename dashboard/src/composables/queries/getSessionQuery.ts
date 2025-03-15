import { env } from "@/env";
import { SessionRequestsSchema } from "@/schema/schema_pb";
import { fromBinary } from "@bufbuild/protobuf";
import { useQuery } from "./useQuery";

export const useGetSessionQuery = () => {
    const query = async (id: string) => {
        const response = await fetch(`${env.baseUrl}/session/${id}`);
        const bytes = await response.bytes();
        const value = fromBinary(SessionRequestsSchema, bytes);
        return value;
    }
    
    return useQuery((id: string) => query(id));
}