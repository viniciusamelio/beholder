import { env } from "@/env";
import  { EnvironmentSessionsSchema } from "@/schema/schema_pb";
import { fromBinary } from "@bufbuild/protobuf";
import { useQuery } from "@pinia/colada";


export const getSessionsQuery = ({ id}  :{
    id: number
}) => {

    return useQuery({
        key: () => ["sessions", id],
        initialData() {
            return undefined;
        },
        query: async () => {
            const response = await fetch(`${env.baseUrl}/environment/${id}/sessions`);
            const bytes = await response.bytes();
            const data = fromBinary(EnvironmentSessionsSchema, bytes);
            console.log(data.sessions);
            return data.sessions;
        },
    });

   
}