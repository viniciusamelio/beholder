<script setup lang="ts">
import { useGetEnvironmentsQuery } from '@/composables/queries/getEnvironmentsQuery';

const { data: environments, isPending } = useGetEnvironmentsQuery();
</script>

<template>
    <aside
        class="sidebar w-64 h-screen sticky top-0 left-0 bg-card border-r-2 border-border flex flex-col text-foreground">
        <div class="p-4 border-b border-border">
            <div class="dropdown w-full">
                <button
                    class="w-full px-3 py-2 text-sm text-left rounded-lg hover:bg-primary hover:text-primary-foreground flex items-center">
                    <i class="ri-command-line mr-2"></i>
                    Acme Inc
                    <i class="ri-arrow-down-s-line ml-auto"></i>
                </button>
            </div>
        </div>

        <!-- Sidebar Content -->
        <nav class="flex-1 overflow-y-auto p-4">
            <div class="space-y-4">

                <div>
                    <h3 class="text-xs font-semibold text-gray-500 mb-2">Environments</h3>
                    <div v-if="isPending">
                        <ul>
                            <li v-for="i in 4">
                                <a href=""
                                    class="sidebar-item flex items-center text-sm px-3 py-2 rounded-lg h-12 w-full animate-pulse bg-background mb-2">
                                    <span class="w-2 h-2 rounded-full bg-blue-500 mr-2 animate-pulse"></span>
                                </a>
                            </li>
                        </ul>
                    </div>
                    <div v-else class="space-y-1">
                        <ul>
                            <li v-for="env in environments">
                                <a href="" class="sidebar-item flex items-center text-sm px-3 py-2 rounded-lg">
                                    <span class="w-2 h-2 rounded-full bg-blue-500 mr-2"></span>
                                    <p v-text="env.name">
                                    </p>
                                </a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </nav>
    </aside>

</template>
