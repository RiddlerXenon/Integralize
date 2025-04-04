import React from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Line } from 'react-chartjs-2';

// Зарегистрируйте все используемые модули
ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

const Chart = ({ diffData }) => {
    if (!diffData) return null;

    const minX = Math.floor(Math.min(...diffData.x));
    const maxX = Math.ceil(Math.max(...diffData.x));
    const minY = Math.floor(Math.min(...diffData.y));
    const maxY = Math.ceil(Math.max(...diffData.y));

    const stepSizeX = (maxX - minX) / 10;
    const stepSizeY = (maxY - minY) / 10;

    const data = {
        labels: diffData.x,
        datasets: [
            {
                data: diffData.y,
                fill: false,
                backgroundColor: 'rgba(75,192,192,0.2)',
                borderColor: 'rgba(75,192,192,1)',
            },
        ],
    };

    const options = {
        scales: {
            x: {
                type: 'linear',
                position: 'bottom',
                beginAtZero: false,
                ticks: {
                    stepSize: stepSizeX,
                    callback: function(value, index, values) {
                        return Number.isInteger(value) ? value : '';
                    }
                },
                title: {
                    display: false,
                    text: 'x',
                },
                min: minX,
                max: maxX,
            },
            y: {
                type: 'linear',
                beginAtZero: false,
                ticks: {
                    stepSize: stepSizeY,
                    callback: function(value, index, values) {
                        return Number.isInteger(value) ? value : '';
                    }
                },
                title: {
                    display: false,
                    text: 'y',
                },
                min: minY,
                max: maxY,
            },
        },
        plugins: {
            legend: {
                display: false, // Отключение отображения легенды
            },
        },
    };

    return <Line data={data} options={options} />;
};

export default Chart;
