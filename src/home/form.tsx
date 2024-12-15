import React, { useState } from 'react';
import { Form, Input, Button, message, InputNumber } from 'antd';

const SampleForm: React.FC = () => {
    const [isSuccess, setIsSuccess] = useState(false);

    const onFinish = async (values: { title: string; priority: number; status: string; worker: string }) => {
        try {
            const response = await fetch("http://localhost:8080/job", {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(values),
            });
            if (response.ok) {
                message.success('Form submitted and data saved successfully!');
                setIsSuccess(true);
            } else {
                throw new Error('Failed to submit form');
            }
        } catch (error) {
            if (error instanceof Error) {
                message.error("Error submitting form: " + error.message);
            } else {
                message.error("An unknown error occurred");
            }
        }
    };

    return (
        <div className="bg-gray-300 w-1/3 justify-center items-center mt-10 py-3">
            <p className="text-slate-600 font-bold">Form</p>

            <Form name="Form" layout="vertical" onFinish={onFinish} className="gap-y-4 w-80">
                <Form.Item label="Title" name="title" rules={[{ required: true, message: 'Please input the title!' }]}>
                    <Input />
                </Form.Item>
                <Form.Item
                    label="Priority"
                    name="priority"
                    className=''
                    rules={[
                        { required: true, message: 'Please input the priority!' },
                        { type: 'number', min: 1, max: 10, message: 'Priority must be between 1 and 10' },
                    ]}
                >
                    <InputNumber min={1} max={10} />
                </Form.Item>
                <Form.Item label="Status" name="status" rules={[{ required: true, message: 'Please input the status!' }]}>
                    <Input />
                </Form.Item>
                <Form.Item label="Worker" name="worker" rules={[{ required: true, message: 'Please input a valid worker!' }]}>
                    <Input />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>

            {isSuccess && (
                <div className="mt-6 bg-gray-50 p-4 border-t-2 border-gray-200">
                    <h3 className="text-xl font-semibold text-center">Form Submitted Successfully</h3>
                </div>
            )}
        </div>
    );
};

export default SampleForm;
