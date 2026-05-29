<template>
  <div class="space-y-lg">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end gap-4">
      <div>
        <h1 class="font-headline-md text-headline-md text-on-surface">Sector Management</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-xs">Orchestrate and partition the industrial hardware landscape.</p>
      </div>
      <button @click="openCreateModal" class="bg-primary hover:bg-primary-fixed-dim text-on-primary font-bold px-lg py-sm rounded-lg transition-all duration-200 flex items-center gap-sm shadow-sm whitespace-nowrap">
        <span class="material-symbols-outlined">add_box</span>
        <span>Define Sector</span>
      </button>
    </div>

    <!-- Sectors Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-md">
      <div v-for="sector in sectors" :key="sector.id" class="glass-panel p-lg rounded-xl flex flex-col justify-between group relative overflow-hidden">
        <div class="absolute inset-0 bg-gradient-to-br from-primary/5 to-transparent opacity-0 group-hover:opacity-100 transition-opacity"></div>
        
        <div class="relative z-10">
          <div class="flex justify-between items-start mb-md">
            <div class="p-3 bg-surface-container-high rounded-lg text-primary">
              <span class="material-symbols-outlined text-[32px]">grid_view</span>
            </div>
            <div class="flex gap-xs">
              <router-link :to="`/sectors/${sector.id}`" class="p-1.5 hover:bg-primary/20 text-on-surface-variant hover:text-primary rounded-lg transition-all" title="View Details">
                <span class="material-symbols-outlined text-[20px]">visibility</span>
              </router-link>
              <button @click="openEditModal(sector)" class="p-1.5 hover:bg-primary/20 text-on-surface-variant hover:text-primary rounded-lg transition-all" title="Edit Partition">
                <span class="material-symbols-outlined text-[20px]">edit</span>
              </button>
              <button @click="deleteSector(sector.id)" class="p-1.5 hover:bg-error/20 text-on-surface-variant hover:text-error rounded-lg transition-all">
                <span class="material-symbols-outlined text-[20px]">delete</span>
              </button>
            </div>
          </div>
          
          <h3 class="font-title-lg text-title-lg text-on-surface mb-xs">{{ sector.name }}</h3>
          <div v-if="sector.parentName" class="flex items-center gap-xs text-primary/80 font-label-caps text-[10px] mb-sm">
            <span class="material-symbols-outlined text-[14px]">subdirectory_arrow_right</span>
            <span>Child of: {{ sector.parentName }}</span>
          </div>
          <p class="font-body-md text-on-surface-variant line-clamp-2 mb-lg">{{ sector.description || 'No description provided for this sector.' }}</p>
        </div>

        <div class="flex items-center justify-between mt-auto pt-md border-t border-white/5 relative z-10">
          <div class="flex items-center gap-xs">
            <span class="material-symbols-outlined text-primary text-[18px]">developer_board</span>
            <span class="font-data-md text-on-surface">{{ sector.deviceCount }} Nodes</span>
          </div>
          <span class="font-label-caps text-[10px] text-on-surface-variant uppercase tracking-widest">Active Partition</span>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="sectors.length === 0 && !isLoading" class="col-span-1 md:col-span-3 glass-panel p-xl rounded-xl text-center">
        <span class="material-symbols-outlined text-on-surface-variant text-5xl mb-md">inventory_2</span>
        <p class="font-title-md text-on-surface">No Industrial Sectors Defined</p>
        <p class="font-body-md text-on-surface-variant mt-xs mb-lg">Partition your network by defining its first operational sector.</p>
        <button @click="openCreateModal" class="text-primary font-bold hover:underline">Provision Initial Sector</button>
      </div>

      <!-- Loading State -->
      <div v-if="isLoading" class="col-span-1 md:col-span-3 py-12 text-center text-on-surface-variant">
        <span class="material-symbols-outlined animate-spin text-3xl">refresh</span>
        <p class="mt-sm font-label-caps">Mapping Sectors...</p>
      </div>
    </div>

    <!-- Modal for Create/Edit -->
    <div v-if="showModal" class="fixed inset-0 z-[60] flex items-center justify-center p-4 sm:p-md">
      <div class="absolute inset-0 bg-surface/80 backdrop-blur-sm" @click="showModal = false"></div>
      <div class="glass-panel w-[calc(100%-2rem)] max-w-md min-w-[300px] rounded-2xl p-lg relative z-10 shadow-2xl border-white/10 overflow-hidden">
        <h2 class="font-headline-sm text-headline-sm text-on-surface mb-md">{{ editingId ? 'Update Sector' : 'Define New Sector' }}</h2>
        
        <form @submit.prevent="saveSector" class="space-y-md">
          <div class="space-y-xs">
            <label class="font-label-caps text-[10px] text-on-surface-variant uppercase">Sector Label</label>
            <input v-model="form.name" required class="w-full bg-surface-container-lowest border border-white/10 rounded-lg px-md py-3 text-on-surface font-body-md focus:border-primary outline-none transition-all" placeholder="e.g. Production Alpha" />
          </div>
          
          <div class="space-y-xs">
            <label class="font-label-caps text-[10px] text-on-surface-variant uppercase">Parent Partition (Optional)</label>
            <select v-model="form.parent_id" class="w-full bg-surface-container-lowest border border-white/10 rounded-lg px-md py-3 text-on-surface font-body-md focus:border-primary outline-none transition-all appearance-none">
              <option :value="null">Root Sector</option>
              <option v-for="s in availableParentSectors" :key="s.id" :value="s.id">
                {{ s.name }}
              </option>
            </select>
          </div>
          
          <div class="space-y-xs">
            <label class="font-label-caps text-[10px] text-on-surface-variant uppercase">Technical Description</label>
            <textarea v-model="form.description" rows="3" class="w-full bg-surface-container-lowest border border-white/10 rounded-lg px-md py-3 text-on-surface font-body-md focus:border-primary outline-none transition-all" placeholder="Operational parameters and purpose..."></textarea>
          </div>

          <div class="flex gap-md pt-md">
            <button type="submit" class="flex-1 bg-primary text-on-primary font-bold py-3 rounded-lg hover:brightness-110 active:scale-95 transition-all" :disabled="isSubmitting">
              {{ isSubmitting ? 'Syncing...' : (editingId ? 'Update Matrix' : 'Confirm Definition') }}
            </button>
            <button @click="showModal = false" type="button" class="px-md border border-white/10 text-on-surface-variant hover:bg-white/5 rounded-lg transition-all">
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { sectorApi } from '../api/sectors'

const sectors = ref([])
const isLoading = ref(false)
const showModal = ref(false)
const isSubmitting = ref(false)
const editingId = ref(null)

const form = ref({
  name: '',
  description: '',
  parent_id: null
})

const availableParentSectors = computed(() => {
  return sectors.value.filter(s => s.id !== editingId.value)
})

const fetchSectors = async () => {
  isLoading.value = true
  try {
    const response = await sectorApi.getSectors(1, 100)
    sectors.value = response.data.map(s => ({
      id: s.ID || s.id,
      name: s.name || s.Name,
      description: s.description || s.Description,
      parent_id: s.parent_id || s.ParentID || null,
      parentName: s.Parent?.name || s.parent?.name || null,
      deviceCount: s.devices?.length || 0
    }))
  } catch (err) {
    console.error('Failed to fetch sectors:', err)
  } finally {
    isLoading.value = false
  }
}

const openCreateModal = () => {
  editingId.value = null
  form.value = { name: '', description: '', parent_id: null }
  showModal.value = true
}

const openEditModal = (sector) => {
  editingId.value = sector.id
  form.value = { 
    name: sector.name, 
    description: sector.description,
    parent_id: sector.parent_id
  }
  showModal.value = true
}

const saveSector = async () => {
  isSubmitting.value = true
  try {
    if (editingId.value) {
      await sectorApi.updateSector(editingId.value, form.value)
    } else {
      await sectorApi.createSector(form.value)
    }
    showModal.value = false
    fetchSectors()
  } catch (err) {
    console.error('Failed to save sector:', err)
    alert('Failed to save sector definition. Check naming uniqueness.')
  } finally {
    isSubmitting.value = false
  }
}

const deleteSector = async (id) => {
  if (!confirm('Permanent deletion will leave associated nodes without a sector assignment. Continue?')) return
  try {
    await sectorApi.deleteSector(id)
    fetchSectors()
  } catch (err) {
    console.error('Failed to delete sector:', err)
    alert('Could not decommission sector. Ensure it is empty or check permissions.')
  }
}

onMounted(() => {
  fetchSectors()
})
</script>

<style scoped>
.glass-panel {
    background: rgba(23, 31, 51, 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.08);
}
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;  
  overflow: hidden;
}
</style>
