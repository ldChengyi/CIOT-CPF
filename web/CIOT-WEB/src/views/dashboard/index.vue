<template>
  <div class="dashboard-container">
    <!-- 数据概览卡片 -->
    <el-row :gutter="20">
      <el-col :span="6" v-for="item in statisticsData" :key="item.title">
        <el-card shadow="hover" class="statistics-card">
          <div class="card-header">
            <span class="title">{{ item.title }}</span>
            <el-icon :class="['icon', item.type]">
              <component :is="item.icon"></component>
            </el-icon>
          </div>
          <div class="card-body">
            <div class="number">{{ item.value }}</div>
            <div class="compare">
              较昨日
              <span :class="item.trend === 'up' ? 'up' : 'down'">
                {{ item.change }}%
                <el-icon>
                  <component :is="item.trend === 'up' ? 'ArrowUp' : 'ArrowDown'"></component>
                </el-icon>
              </span>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 设备状态图表 -->
    <el-row :gutter="20" class="chart-row">
      <el-col :span="16">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>设备在线趋势</span>
              <el-radio-group v-model="timeRange" size="small">
                <el-radio-button label="week">近一周</el-radio-button>
                <el-radio-button label="month">近一月</el-radio-button>
              </el-radio-group>
            </div>
          </template>
          <div class="chart" ref="lineChartRef"></div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>设备状态分布</span>
            </div>
          </template>
          <div class="chart" ref="pieChartRef"></div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 最近活动 -->
    <el-card shadow="hover" class="activity-card">
      <template #header>
        <div class="card-header">
          <span>最近活动</span>
          <el-button type="primary" link>查看更多</el-button>
        </div>
      </template>
      <el-timeline>
        <el-timeline-item
          v-for="(activity, index) in activities"
          :key="index"
          :type="activity.type"
          :timestamp="activity.timestamp"
        >
          {{ activity.content }}
        </el-timeline-item>
      </el-timeline>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as echarts from 'echarts'
import { Monitor, Connection, Box, Link } from '@element-plus/icons-vue'

// 统计数据
const statisticsData = ref([
  {
    title: '设备总数',
    value: '2,645',
    change: '15',
    trend: 'up',
    icon: 'Monitor',
    type: 'primary'
  },
  {
    title: '在线设备',
    value: '1,578',
    change: '8',
    trend: 'up',
    icon: 'Connection',
    type: 'success'
  },
  {
    title: '产品数量',
    value: '36',
    change: '2',
    trend: 'down',
    icon: 'Box',
    type: 'warning'
  },
  {
    title: '告警设备',
    value: '12',
    change: '50',
    trend: 'up',
    icon: 'Link',
    type: 'danger'
  }
])

// 时间范围选择
const timeRange = ref('week')

// 最近活动
const activities = ref([
  {
    content: '设备 "温度传感器_01" 触发高温告警',
    timestamp: '2024-01-20 10:30:00',
    type: 'warning'
  },
  {
    content: '新增设备 "压力传感器_02" 成功接入',
    timestamp: '2024-01-20 09:15:00',
    type: 'success'
  },
  {
    content: '设备 "湿度传感器_03" 离线',
    timestamp: '2024-01-20 08:45:00',
    type: 'danger'
  },
  {
    content: '更新设备固件版本至 v2.1.0',
    timestamp: '2024-01-20 08:00:00',
    type: 'primary'
  }
])

// 图表实例
let lineChart = null
let pieChart = null

// 初始化折线图
const initLineChart = (dom) => {
  lineChart = echarts.init(dom)
  const option = {
    tooltip: {
      trigger: 'axis'
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        name: '在线设备',
        type: 'line',
        smooth: true,
        data: [1200, 1320, 1500, 1480, 1650, 1800, 1750],
        areaStyle: {
          opacity: 0.1
        },
        lineStyle: {
          width: 3
        }
      }
    ]
  }
  lineChart.setOption(option)
}

// 初始化饼图
const initPieChart = (dom) => {
  pieChart = echarts.init(dom)
  const option = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'vertical',
      left: 'left'
    },
    series: [
      {
        type: 'pie',
        radius: '60%',
        data: [
          { value: 1578, name: '在线' },
          { value: 856, name: '离线' },
          { value: 211, name: '故障' }
        ],
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)'
          }
        }
      }
    ]
  }
  pieChart.setOption(option)
}

// 监听窗口大小变化
const handleResize = () => {
  lineChart?.resize()
  pieChart?.resize()
}

onMounted(() => {
  const lineChartDom = document.querySelector('.chart-row .chart')
  const pieChartDom = document.querySelector('.chart-row .el-col:last-child .chart')
  
  initLineChart(lineChartDom)
  initPieChart(pieChartDom)
  
  window.addEventListener('resize', handleResize)
})
</script>

<style lang="scss" scoped>
.dashboard-container {
  padding: 20px;
  
  .statistics-card {
    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
      
      .title {
        font-size: 16px;
        color: #606266;
      }
      
      .icon {
        font-size: 24px;
        
        &.primary { color: #409EFF; }
        &.success { color: #67C23A; }
        &.warning { color: #E6A23C; }
        &.danger { color: #F56C6C; }
      }
    }
    
    .card-body {
      .number {
        font-size: 24px;
        font-weight: bold;
        margin-bottom: 10px;
      }
      
      .compare {
        font-size: 14px;
        color: #909399;
        
        .up {
          color: #67C23A;
        }
        
        .down {
          color: #F56C6C;
        }
      }
    }
  }
  
  .chart-row {
    margin-top: 20px;
    
    .chart {
      height: 350px;
    }
    
    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
  }
  
  .activity-card {
    margin-top: 20px;
    
    .card-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
  }
}

// 行间距
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}
</style>