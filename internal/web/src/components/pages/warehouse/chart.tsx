'use client';

import * as React from 'react';
import { Label, Pie, PieChart } from 'recharts';

import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from '@/components/ui/card';
import {
    ChartConfig,
    ChartContainer,
    ChartTooltip,
    ChartTooltipContent,
} from '@/components/ui/chart';
import { Warehouse, WarehouseItem } from '@/lib/zod.schema';

export const description = 'A donut chart with text';

const ColorMap = [
    'hsl(var(--chart-1))',
    'hsl(var(--chart-2))',
    'hsl(var(--chart-3))',
    'hsl(var(--chart-4))',
    'hsl(var(--chart-5))',
];

function getChartData(
    warehouse: Warehouse,
    warehouseItems: WarehouseItem[]
): { config: ChartConfig; data: unknown[] } {
    const sortList = warehouseItems.sort((v1, v2) => v2.capacity - v1.capacity);

    let otherCapacity = 0;
    let usedCapacity = 0;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const conf: any = {
        capacity: {
            label: 'capacity',
        },
    };
    const data: unknown[] = [];

    sortList.forEach((vl, ind) => {
        if (ind < 3) {
            conf[vl.id] = {
                label: vl.rice_name,
                color: ColorMap[ind],
            };

            data.push({
                capacity: vl.capacity,
                name: vl.rice_name,
                fill: ColorMap[ind],
            });
        } else {
            otherCapacity += vl.capacity;
        }
        usedCapacity += vl.capacity;
    });
    data.push({
        capacity: otherCapacity,
        name: 'Other',
        fill: ColorMap[3],
    });
    conf['Other'] = {
        label: 'Other',
        color: ColorMap[3],
    };

    data.push({
        capacity: warehouse.capacity - usedCapacity,
        name: 'Free',
        fill: ColorMap[4],
    });
    conf['Free'] = {
        label: 'Free',
        color: ColorMap[4],
    };

    return {
        config: conf,
        data: data,
    };
}

type Props = {
    warehouse: Warehouse;
    warehouseItems: WarehouseItem[];
};

export function CapacityPieChart({ warehouse, warehouseItems }: Props) {
    const totalUsedCapacity = React.useMemo(() => {
        return warehouseItems.reduce((p, v) => p + v.capacity, 0);
    }, [warehouseItems]);

    const { config, data } = React.useMemo(
        () => getChartData(warehouse, warehouseItems),
        [warehouse, warehouseItems]
    );

    return (
        <Card className="flex flex-col ">
            <CardHeader className="items-center pb-0 ">
                <CardTitle>Warehouse inventory</CardTitle>
                <CardDescription>{warehouse.name}</CardDescription>
            </CardHeader>
            <CardContent className="flex-1 pb-1">
                <ChartContainer
                    config={config}
                    className="mx-auto aspect-square max-h-[250px] -my-3"
                >
                    <PieChart>
                        <ChartTooltip
                            cursor={false}
                            content={
                                <ChartTooltipContent
                                    className="bg-white"
                                    hideLabel
                                />
                            }
                        />
                        <Pie
                            data={data}
                            dataKey="capacity"
                            nameKey="name"
                            innerRadius={60}
                            strokeWidth={5}
                        >
                            <Label
                                content={({ viewBox }) => {
                                    if (
                                        viewBox &&
                                        'cx' in viewBox &&
                                        'cy' in viewBox
                                    ) {
                                        return (
                                            <text
                                                x={viewBox.cx}
                                                y={viewBox.cy}
                                                textAnchor="middle"
                                                dominantBaseline="middle"
                                            >
                                                <tspan
                                                    x={viewBox.cx}
                                                    y={viewBox.cy}
                                                    className="fill-foreground text-3xl font-bold"
                                                >
                                                    {totalUsedCapacity.toLocaleString()}
                                                </tspan>
                                                <tspan
                                                    x={viewBox.cx}
                                                    y={(viewBox.cy || 0) + 24}
                                                    className="fill-muted-foreground"
                                                >
                                                    Used capacity
                                                </tspan>
                                            </text>
                                        );
                                    }
                                }}
                            />
                        </Pie>
                    </PieChart>
                </ChartContainer>
            </CardContent>
        </Card>
    );
}
