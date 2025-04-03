import { ref, type Ref } from "vue";

export const useQuery = <T>(
    query:(args: any)=> Promise<T>
) => {
    const isPending = ref(false);
    const isError = ref(false);
    const isSuccess = ref(false);
    const isIdle = ref(true);
    const data = ref<T | null>(null);

    const wrappedQuery = (args: any) => {
        isIdle.value = false;
        isPending.value = true;
        query(args).then((value) => {
            isSuccess.value = true;
            isError.value = false;
            data.value = value;
        }).catch(() => {
            isSuccess.value = true;
            isError.value = true;
        }).finally(() => {
            isPending.value = false;
        })
    }

    return {
        query: wrappedQuery,
        refetch: wrappedQuery,
        isError,
        isIdle,
        isSuccess,
        isPending,
        data
    }
}