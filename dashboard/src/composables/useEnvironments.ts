
import { defineStore } from "pinia";
import { computed, ref, watch } from "vue";
import { useGetSessionsQuery } from "./queries/getSessionsQuery";
import type { Environment } from "@/types";
import { useGetEnvironmentsQuery } from "./queries/getEnvironmentsQuery";

export const useEnvironments = defineStore("environments", () => {
    const selectedEnvironment = ref<Environment | null>(null);
    const id = computed(() => selectedEnvironment.value?.id ?? 0);

    return {
        selectEnvironment : (environment: Environment) => selectedEnvironment.value = environment,  
        selectedEnvironment,
        id,
        environments : () => useGetEnvironmentsQuery(),
    }
});