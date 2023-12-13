import React, { useState, useEffect } from "react";
import ReactApexChart from "react-apexcharts";
import { EventsOn } from "../../wailsjs/runtime/runtime";

const PieChart = () => {
  const [series, setSeries] = useState([]);
  const [options] = useState({
    chart: {
      width: 180,
      type: "donut",
    },
    dataLabels: {
      enabled: false,
    },
    legend: {
      show: false,
    },
    title: {
      text: "CPU Usage", // Title text for the chart
      align: "center",
      style: {
        fontSize: "14px",
        color: "#ffffff",
      },
    },
    plotOptions: {
      pie: {
        donut: {
          labels: {
            show: true,
            name: {
              offsetY: -3,
              show: true,
              color: "#888",
              fontSize: "17px",
            },
            value: {
              offsetY: 3,
              color: "#fff",
              fontSize: "14px",
              show: true,
            },
            total: {
              show: true,
              label: "Usage",
              color: "#fff",
              fontSize: "14px",
              formatter: function (w) {
                return w.globals.seriesTotals.reduce((a, b) => {
                  return 100 - b + "%";
                }, 0);
              },
            },
          },
        },
      },
    },
    responsive: [
      {
        breakpoint: 280,
        options: {
          chart: {
            width: 200,
          },
          legend: {
            show: true,
          },
        },
      },
    ],
  });

  useEffect(() => {
    EventsOn("cpudata", (data) => {
      const p = data["usage"];
      const r = 100 - p;
      setSeries([p, r]);
    });
  }, []);

  return (
    <div>
      <div className="chart-wrap">
        <div id="chart">
          <ReactApexChart
            options={options}
            series={series}
            type="donut"
            width={200}
          />
        </div>
      </div>
    </div>
  );
};

export default PieChart;
