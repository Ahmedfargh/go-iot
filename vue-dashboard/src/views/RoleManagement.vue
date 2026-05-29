<template>
  <div class="space-y-lg px-margin-mobile md:px-0">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end gap-4">
      <div>
        <h1 class="font-headline-md text-headline-md text-on-surface">Access Tiers & Roles</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-xs">Define system permissions and operational scope for technical personnel.</p>
      </div>
      <button @click="showRoleModal = true" class="bg-primary hover:bg-primary-fixed-dim text-on-primary font-bold px-lg py-sm rounded-lg transition-all duration-200 flex items-center gap-sm shadow-sm whitespace-nowrap">
        <span class="material-symbols-outlined">add_moderator</span>
        <span>Define New Role</span>
      </button>
    </div>

    <!-- Roles Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-md">
      <div v-for="role in roles" :key="role.name" class="glass-panel p-lg rounded-xl flex flex-col gap-md relative overflow-hidden group">
        <div class="absolute top-0 right-0 w-16 h-16 bg-gradient-to-bl from-primary/5 to-transparent transition-all group-hover:from-primary/15"></div>
        
        <div class="flex items-center gap-md">
          <div class="w-12 h-12 rounded-xl bg-surface-container-high flex items-center justify-center text-primary">
            <span class="material-symbols-outlined text-[28px]">{{ getRoleIcon(role.name) }}</span>
          </div>
          <div>
            <h3 class="font-title-sm text-title-sm text-on-surface">{{ role.name }}</h3>
            <p class="font-label-caps text-[10px] text-primary">{{ role.permissions?.length || 0 }} PERSISTED PERMISSIONS</p>
          </div>
        </div>

        <p class="text-body-md text-on-surface-variant line-clamp-2 min-h-[3rem]">{{ role.description || 'Standard access tier for technical nodes.' }}</p>

        <div class="flex flex-wrap gap-xs">
          <span v-for="permission in role.permissions?.slice(0, 3)" :key="permission" class="px-2 py-0.5 bg-surface-container-highest border border-white/5 rounded font-label-caps text-[9px] text-on-surface-variant uppercase">
            {{ permission }}
          </span>
          <span v-if="role.permissions?.length > 3" class="px-2 py-0.5 bg-surface-container-highest border border-white/5 rounded font-label-caps text-[9px] text-primary">
            +{{ role.permissions.length - 3 }} MORE
          </span>
        </div>

        <div class="mt-auto pt-md border-t border-white/5 flex justify-between items-center">
          <button @click="editPermissions(role)" class="text-label-caps text-primary hover:brightness-110 font-bold transition-all text-[11px] flex items-center gap-xs">
            <span class="material-symbols-outlined text-[16px]">sync_alt</span>
            SYNC PERMISSIONS
          </button>
          <span class="material-symbols-outlined text-on-surface-variant/30 text-[20px]">verified_user</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { roleApi } from '../api/roles'

const roles = ref([])
const isLoading = ref(false)
const showRoleModal = ref(false)

const getRoleIcon = (name) => {
  const n = name.toLowerCase()
  if (n === 'root') return 'terminal'
  if (n === 'security') return 'security'
  if (n === 'network') return 'hub'
  if (n === 'operator') return 'manage_accounts'
  return 'settings_input_composite'
}

const fetchRoles = async () => {
  isLoading.value = true
  try {
    const data = await roleApi.getRoles()
    roles.value = data.map(r => ({
      name: r.Name,
      description: r.Description,
      permissions: r.Permissions || []
    }))
  } catch (err) {
    console.error('Failed to load roles repository:', err)
  } finally {
    isLoading.value = false
  }
}

const editPermissions = (role) => {
  // Logic to open a permission sync modal
  console.log('Syncing permissions for:', role.name)
}

onMounted(() => {
  fetchRoles()
})
</script>

<style scoped>
.glass-panel {
    background: rgba(23, 31, 51, 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.08);
}
</style>
