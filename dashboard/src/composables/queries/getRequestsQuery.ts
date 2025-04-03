import { env } from "@/env";
import { EnvironmentRequestsSchema, type Request } from "@/schema/schema_pb";
import { fromBinary } from "@bufbuild/protobuf";
import { useQuery } from "./useQuery";


export const useGetRequestsQuery = () => {
    const query = async (id: string) => {
        const response = await fetch(`${env.baseUrl}/environment/${id}/requests`);
        const bytes = await response.bytes();
        const value = fromBinary(EnvironmentRequestsSchema, bytes);
        return value.requests;
    }

    return useQuery((id:string) => query(id));
}