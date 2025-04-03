<script setup lang="ts">
const { method } = defineProps<{ method: string }>();


const getMethodColor = (method: string) => {
    switch (method) {
      case "GET":
        return "bg-emerald-500/10 text-emerald-500  border-emerald-500/20"
      case "POST":
        return "bg-blue-500/10 text-blue-500 border-blue-500/20"
      case "PUT":
        return "bg-amber-500/10 text-amber-500 border-amber-500/20"
      case "DELETE":
        return "bg-red-500/10 text-red-500 border-red-500/20"
      case "PATCH":
        return "bg-purple-500/10 text-purple-500 border-purple-500/20"
      default:
        return "bg-gray-500/10 text-gray-400 border-gray-500/20"
    }
}

</script>

<template>
    <div class="mr-2 font-mono text-xs px-2 py-0.5 border rounded-xl flex items-center justify-center border-border" :class="getMethodColor(method)">
        <p v-text="method"></p>
    </div>
</template>