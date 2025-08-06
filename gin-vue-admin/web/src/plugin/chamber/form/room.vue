
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="房间名称:" prop="roomName">
          <el-input v-model="formData.roomName" :clearable="true"  placeholder="请输入房间名称" />
       </el-form-item>
        <el-form-item label="房主用户ID，关联 user.id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="房间类型(0-公开房,1-私密房):" prop="roomType">
           <el-select v-model="formData.roomType" placeholder="请选择房间类型(0-公开房,1-私密房)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in roomOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="房间状态(0-正常,1-已解散,2-全局禁言):" prop="status">
           <el-select v-model="formData.status" placeholder="请选择房间状态(0-正常,1-已解散,2-全局禁言)" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in statuseOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="房间公告:" prop="announcement">
          <el-input v-model="formData.announcement" :clearable="true"  placeholder="请输入房间公告" />
       </el-form-item>
        <el-form-item label="解散时间:" prop="closedAt">
          <el-date-picker v-model="formData.closedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="最后活跃时间:" prop="lastActiveTime">
          <el-date-picker v-model="formData.lastActiveTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="房间人数:" prop="fkMemberRoom">
          <el-input v-model.number="formData.fkMemberRoom" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="房间封面图:" prop="image">
          <el-input v-model="formData.image" :clearable="true"  placeholder="请输入房间封面图" />
       </el-form-item>
        <el-form-item label="关联的标签ID（一个房间只能选一个）:" prop="tagId">
          <el-input v-model.number="formData.tagId" :clearable="true" placeholder="请输入" />
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
  createRoom,
  updateRoom,
  findRoom
} from '@/plugin/chamber/api/room'

defineOptions({
    name: 'RoomForm'
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
const roomOptions = ref([])
const statuseOptions = ref([])
const formData = ref({
            roomName: '',
            userId: undefined,
            roomType: '',
            status: '',
            announcement: '',
            closedAt: new Date(),
            lastActiveTime: new Date(),
            fkMemberRoom: undefined,
            image: '',
            tagId: undefined,
        })
// 验证规则
const rule = reactive({
               roomName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roomType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findRoom({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    roomOptions.value = await getDictFunc('room')
    statuseOptions.value = await getDictFunc('statuse')
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
               res = await createRoom(formData.value)
               break
             case 'update':
               res = await updateRoom(formData.value)
               break
             default:
               res = await createRoom(formData.value)
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
