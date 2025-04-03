import { env } from "@/env";
import { EnvironmentSessionsSchema, type Session } from "@/schema/schema_pb";
import { fromBinary } from "@bufbuild/protobuf";
import { useQuery } from "./useQuery";


export const useGetSessionsQuery = () => {
    const query = async (id: string) => {
        const response = await fetch(`${env.baseUrl}/environment/${id}/sessions`);
        const bytes = await response.bytes();
        const value = fromBinary(EnvironmentSessionsSchema, bytes);
        return value.sessions;
    }
    return useQuery((id: string) => query(id))
}