import React from 'react';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';

const CustomTabs = ({ renderIntegralForm, renderDifferentialForm, renderPredatorPreyForm }) => {
    return (
        <Tabs>
            <TabList>
                <Tab>Интегралы</Tab>
                <Tab>Дифференциальные Уравнения</Tab>
                <Tab>Модель Хищник-Жертва</Tab>
            </TabList>

            <TabPanel>
                <h2>Интегральные Уравнения</h2>
                {renderIntegralForm()}
            </TabPanel>

            <TabPanel>
                <h2>Дифференциальные Уравнения</h2>
                {renderDifferentialForm()}
            </TabPanel>

            <TabPanel>
                <h2>Модель Хищник-Жертва</h2>
                {renderPredatorPreyForm()}
            </TabPanel>
        </Tabs>
    );
};

export default CustomTabs;
