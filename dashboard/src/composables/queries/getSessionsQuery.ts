import { env } from "@/env";
import  { EnvironmentSessionsSchema } from "@/schema/schema_pb";
import { fromBinary } from "@bufbuild/protobuf";
import { useQuery } from "@pinia/colada";
import { useEnvironments } from "../useEnvironments";


export const getSessionsQuery = () => {
    const environment = useEnvironments();
    const id = environment.id;
    return useQuery({
        key: () => ["sessions", id],
        initialData() {
            return undefined;
        },
        enabled: id > 0,
        refetchOnReconnect: true,
        query: async () => {
            const response = await fetch(`${env.baseUrl}/environment/${id}/sessions`);
            const bytes = await response.bytes();
            const data = fromBinary(EnvironmentSessionsSchema, bytes);
            console.log(data.sessions);
            return data.sessions;
        },
    });

   
}