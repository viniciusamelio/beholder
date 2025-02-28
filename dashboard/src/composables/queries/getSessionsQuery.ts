import { env } from "@/env";
import { EnvironmentSessionsSchema } from "@/schema/schema_pb";
import { fromBinary } from "@bufbuild/protobuf";
import { ref } from "vue";


export const useGetSessionsQuery =  <T>() => {
    const isPending = ref(false);
    const isError = ref(false);
    const isSuccess = ref(false);
    const isIdle = ref(true);
    const data = ref<T | null>(null);
    
    const query = async (id:string) => {
        try {
            isIdle.value = false;
            isPending.value = true;
            const response = await fetch(`${env.baseUrl}/environment/${id}/sessions`);
            const bytes = await response.bytes();
            const value = fromBinary(EnvironmentSessionsSchema, bytes);
            console.log(value.sessions);
            isSuccess.value = true;
            isError.value = false;
            data.value = value.sessions;
            return value.sessions;
        } catch (error) {
            isSuccess.value = true;
            isError.value = true;
        } finally {
            isPending.value = false;
        }
    }

    return {
        query,
        refetch: (id:string)=>query(id),
        isError,
        isIdle,
        isSuccess,
        isPending,
        data
    }
}