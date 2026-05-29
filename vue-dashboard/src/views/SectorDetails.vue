<template>
  <div class="space-y-lg">
    <!-- Breadcrumbs & Header -->
    <header class="flex flex-col gap-md">
      <div class="flex items-center gap-sm">
        <router-link to="/sectors" class="p-2 -ml-2 text-on-surface-variant hover:bg-surface-container-high rounded-lg transition-all active:scale-95">
          <span class="material-symbols-outlined">arrow_back</span>
        </router-link>
        <div class="flex items-center gap-xs font-label-caps text-[10px] text-on-surface-variant uppercase tracking-widest">
          <router-link to="/sectors" class="hover:text-primary">Sectors</router-link>
          <span class="material-symbols-outlined text-[14px]">chevron_right</span>
          <span class="text-on-surface">Partition Details</span>
        </div>
      </div>
      
      <div class="flex flex-col md:flex-row justify-between items-start md:items-end gap-4">
        <div v-if="sector">
          <div class="flex items-center gap-md mb-xs">
            <h1 class="font-headline-md text-headline-md text-on-surface">{{ sector.name }}</h1>
            <span class="px-3 py-1 bg-primary/10 text-primary border border-primary/20 rounded-full font-label-caps text-[10px] uppercase tracking-tighter">
              {{ sector.devices?.length || 0 }} NODES ACTIVE
            </span>
          </div>
          <p class="font-body-md text-on-surface-variant max-w-2xl">{{ sector.description || 'No technical specifications provided for this industrial sector.' }}</p>
        </div>
        <div v-else-if="isLoading" class="animate-pulse space-y-sm">
          <div class="h-8 w-64 bg-surface-container-high rounded-lg"></div>
          <div class="h-4 w-96 bg-surface-container rounded-lg"></div>
        </div>
        
        <div class="flex items-center gap-sm w-full md:w-auto">
          <div v-if="sector?.parent" class="glass-panel px-md py-sm flex items-center gap-md rounded-xl border-primary/20 bg-primary/5">
            <span class="material-symbols-outlined text-primary">lan</span>
            <div class="flex flex-col">
              <span class="font-label-caps text-[9px] text-on-surface-variant uppercase">Parent Partition</span>
              <router-link :to="`/sectors/${sector.parent.ID}`" class="font-title-sm text-title-sm text-on-surface hover:text-primary transition-colors">
                {{ sector.parent.name }}
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-lg">
      <!-- Left Column: Child Sectors -->
      <div class="lg:col-span-1 space-y-lg">
        <section class="glass-panel p-md rounded-xl h-full">
          <div class="flex items-center justify-between mb-lg">
            <h2 class="font-title-md text-on-surface-variant uppercase tracking-widest flex items-center gap-sm">
              <span class="material-symbols-outlined text-primary">account_tree</span>
              Sub-Partitions
            </h2>
            <span class="font-data-md text-on-surface-variant">{{ sector?.children?.length || 0 }}</span>
          </div>
          
          <div class="space-y-md">
            <div v-for="child in sector?.children" :key="child.ID" class="p-md bg-surface-container-low rounded-xl border border-white/5 hover:border-primary/30 transition-all group">
              <router-link :to="`/sectors/${child.ID}`" class="block">
                <div class="flex justify-between items-center mb-xs">
                  <span class="font-title-sm text-on-surface group-hover:text-primary transition-colors">{{ child.name }}</span>
                  <span class="material-symbols-outlined text-on-surface-variant group-hover:translate-x-1 transition-transform">chevron_right</span>
                </div>
                <p class="text-[12px] text-on-surface-variant line-clamp-1">{{ child.description || 'No description provided.' }}</p>
              </router-link>
            </div>
            
            <div v-if="!sector?.children?.length && !isLoading" class="py-12 text-center text-on-surface-variant border-2 border-dashed border-white/5 rounded-2xl">
              <span class="material-symbols-outlined text-4xl mb-sm opacity-20">grid_view</span>
              <p class="font-label-caps">No sub-partitions defined</p>
            </div>
          </div>
        </section>
      </div>

      <!-- Right Column: Devices -->
      <div class="lg:col-span-2 space-y-lg">
        <section class="glass-panel rounded-xl overflow-hidden min-h-[400px]">
          <div class="px-md py-md bg-white/5 border-b border-white/10 flex justify-between items-center">
            <h2 class="font-title-md text-on-surface-variant uppercase tracking-widest flex items-center gap-sm">
              <span class="material-symbols-outlined text-primary">precision_manufacturing</span>
              Deployed Hardware Nodes
            </h2>
          </div>
          
          <div class="overflow-x-auto">
            <table class="w-full border-collapse">
              <thead>
                <tr class="bg-surface-container-high/50 text-left border-b border-white/10">
                  <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase">Identifier</th>
                  <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-center">Status</th>
                  <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-center">Temp</th>
                  <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-center">Hum</th>
                  <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-right">Actions</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/5">
                <tr v-for="device in sector?.devices" :key="device.ID" class="hover:bg-white/5 transition-colors group">
                  <td class="px-md py-4">
                    <span class="font-data-md text-on-surface font-bold tracking-wider">{{ device.identifier }}</span>
                  </td>
                  <td class="px-md py-4 text-center">
                    <div class="flex justify-center">
                      <div class="flex items-center gap-xs bg-surface-container-high px-sm py-[2px] rounded-full border border-white/5">
                        <span class="w-1.5 h-1.5 rounded-full" 
                              :class="device.status === 'online' ? 'bg-primary' : 'bg-error'"></span>
                        <span class="font-label-caps text-[10px] uppercase text-on-surface-variant">{{ device.status }}</span>
                      </div>
                    </div>
                  </td>
                  <td class="px-md py-4 text-center">
                    <span class="font-data-md text-on-surface">{{ device.temperature }}°C</span>
                  </td>
                  <td class="px-md py-4 text-center">
                    <span class="font-data-md text-on-surface">{{ device.humidity }}%</span>
                  </td>
                  <td class="px-md py-4 text-right">
                    <router-link :to="`/devices/configure/${device.ID}`" class="p-2 hover:bg-primary/20 text-on-surface-variant hover:text-primary rounded-lg transition-all active:scale-95 inline-block">
                      <span class="material-symbols-outlined text-[18px]">settings</span>
                    </router-link>
                  </td>
                </tr>
                <tr v-if="!sector?.devices?.length && !isLoading">
                  <td colspan="5" class="py-24 text-center text-on-surface-variant">
                    <span class="material-symbols-outlined text-5xl mb-sm opacity-20">developer_board_off</span>
                    <p class="font-label-caps uppercase tracking-tighter">No hardware nodes deployed in this sector</p>
                    <router-link to="/devices/provision" class="text-primary mt-md inline-block hover:underline font-title-sm">Provision Initial Node</router-link>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { sectorApi } from '../api/sectors'

const route = useRoute()
const sector = ref(null)
const isLoading = ref(false)

const fetchDetails = async () => {
    isLoading.value = true
    try {
        sector.value = await sectorApi.getSector(route.params.id)
    } catch (err) {
        console.error('Failed to resolve partition details:', err)
    } finally {
        isLoading.value = false
    }
}

onMounted(() => {
    fetchDetails()
})

// Refetch if ID changes (when clicking on child/parent links)
watch(() => route.params.id, (newId) => {
    if (newId) fetchDetails()
})
</script>

<style scoped>
.glass-panel {
    background: rgba(23, 31, 51, 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.08);
}
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;  
  overflow: hidden;
  line-clamp: 1;
}
</style>
