<script setup lang="ts">
import type { Request } from '@/schema/schema_pb';

type Position = "first" | "last" | "middle"
const { request } = defineProps < { request: Request, position: Position } > ();
</script>

<template>
    <div class="rounded-lg bg-card px-3 py-2 cursor-pointer">
        <span class="text-sm text-card-foreground" v-text="request.name"></span>
    </div>
    <!-- <span v-if="['middle', 'last'].includes(position)" class="relative right-4 top-3 rounded-full h-5 w-5 bg-card"/> -->
    <!-- <span v-if="['middle', 'first'].includes(position)" class="relative left-4 top-3 rounded-full h-5 w-5 bg-card"/> -->
    
</template>