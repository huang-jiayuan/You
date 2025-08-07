
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="房间id:" prop="roomId">
    <el-input v-model.number="formData.roomId" :clearable="true" placeholder="请输入房间id" />
</el-form-item>
        <el-form-item label="房间类型:" prop="roomType">
    <el-select v-model="formData.roomType" placeholder="请选择房间类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in room_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="房间标签:" prop="roomTags">
    <el-input v-model="formData.roomTags" :clearable="true" placeholder="请输入房间标签" />
</el-form-item>
        <el-form-item label="房间状态:" prop="roomStatus">
    <el-select v-model="formData.roomStatus" placeholder="请选择房间状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in rom_statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="房主:" prop="roomHost">
    <el-input v-model="formData.roomHost" :clearable="true" placeholder="请输入房主" />
</el-form-item>
        <el-form-item label="房间人数:" prop="roomNum">
    <el-input v-model.number="formData.roomNum" :clearable="true" placeholder="请输入房间人数" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createHostRoom,
  updateHostRoom,
  findHostRoom
} from '@/api/hot_room/hostRoom'

defineOptions({
    name: 'HostRoomForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const room_typeOptions = ref([])
const rom_statusOptions = ref([])
const formData = ref({
            roomId: undefined,
            roomType: '',
            roomTags: '',
            roomStatus: '',
            roomHost: '',
            roomNum: undefined,
        })
// 验证规则
const rule = reactive({
               roomId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomTags : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomHost : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomNum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findHostRoom({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    room_typeOptions.value = await getDictFunc('room_type')
    rom_statusOptions.value = await getDictFunc('rom_status')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createHostRoom(formData.value)
               break
             case 'update':
               res = await updateHostRoom(formData.value)
               break
             default:
               res = await createHostRoom(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
