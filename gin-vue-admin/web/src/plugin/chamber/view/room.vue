
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="w-[380px]"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
       </el-form-item>
      
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="chamber_Room" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="chamber_Room" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="chamber_Room" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="房间名称" prop="roomName" width="120" />

            <el-table-column align="left" label="房主用户ID，关联 user.id" prop="userId" width="120" />

            <el-table-column align="left" label="房间类型(0-公开房,1-私密房)" prop="roomType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.roomType,roomOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="房间状态(0-正常,1-已解散,2-全局禁言)" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,statuseOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="房间公告" prop="announcement" width="120" />

            <el-table-column align="left" label="解散时间" prop="closedAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.closedAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="最后活跃时间" prop="lastActiveTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastActiveTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="房间人数" prop="fkMemberRoom" width="120" />

            <el-table-column align="left" label="房间封面图" prop="image" width="120" />

            <el-table-column align="left" label="关联的标签ID（一个房间只能选一个）" prop="tagId" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateRoomFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
             <el-form-item label="房间名称:" prop="roomName">
    <el-input v-model="formData.roomName" :clearable="true" placeholder="请输入房间名称" />
</el-form-item>
             <el-form-item label="房主用户ID，关联 user.id:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入房主用户ID，关联 user.id" />
</el-form-item>
             <el-form-item label="房间类型(0-公开房,1-私密房):" prop="roomType">
    <el-select v-model="formData.roomType" placeholder="请选择房间类型(0-公开房,1-私密房)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in roomOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="房间状态(0-正常,1-已解散,2-全局禁言):" prop="status">
    <el-select v-model="formData.status" placeholder="请选择房间状态(0-正常,1-已解散,2-全局禁言)" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in statuseOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="房间公告:" prop="announcement">
    <el-input v-model="formData.announcement" :clearable="true" placeholder="请输入房间公告" />
</el-form-item>
             <el-form-item label="解散时间:" prop="closedAt">
    <el-date-picker v-model="formData.closedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="最后活跃时间:" prop="lastActiveTime">
    <el-date-picker v-model="formData.lastActiveTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="房间人数:" prop="fkMemberRoom">
    <el-input v-model.number="formData.fkMemberRoom" :clearable="true" placeholder="请输入房间人数" />
</el-form-item>
             <el-form-item label="房间封面图:" prop="image">
    <el-input v-model="formData.image" :clearable="true" placeholder="请输入房间封面图" />
</el-form-item>
             <el-form-item label="关联的标签ID（一个房间只能选一个）:" prop="tagId">
    <el-input v-model.number="formData.tagId" :clearable="true" placeholder="请输入关联的标签ID（一个房间只能选一个）" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="房间名称">
    {{ detailFrom.roomName }}
</el-descriptions-item>
                 <el-descriptions-item label="房主用户ID，关联 user.id">
    {{ detailFrom.userId }}
</el-descriptions-item>
                 <el-descriptions-item label="房间类型(0-公开房,1-私密房)">
    {{ detailFrom.roomType }}
</el-descriptions-item>
                 <el-descriptions-item label="房间状态(0-正常,1-已解散,2-全局禁言)">
    {{ detailFrom.status }}
</el-descriptions-item>
                 <el-descriptions-item label="房间公告">
    {{ detailFrom.announcement }}
</el-descriptions-item>
                 <el-descriptions-item label="解散时间">
    {{ detailFrom.closedAt }}
</el-descriptions-item>
                 <el-descriptions-item label="最后活跃时间">
    {{ detailFrom.lastActiveTime }}
</el-descriptions-item>
                 <el-descriptions-item label="房间人数">
    {{ detailFrom.fkMemberRoom }}
</el-descriptions-item>
                 <el-descriptions-item label="房间封面图">
    {{ detailFrom.image }}
</el-descriptions-item>
                 <el-descriptions-item label="关联的标签ID（一个房间只能选一个）">
    {{ detailFrom.tagId }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createRoom,
  deleteRoom,
  deleteRoomByIds,
  updateRoom,
  findRoom,
  getRoomList
} from '@/plugin/chamber/api/room'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'Room'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               roomType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getRoomList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    roomOptions.value = await getDictFunc('room')
    statuseOptions.value = await getDictFunc('statuse')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteRoomFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteRoomByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateRoomFunc = async(row) => {
    const res = await findRoom({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteRoomFunc = async (row) => {
    const res = await deleteRoom({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findRoom({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
