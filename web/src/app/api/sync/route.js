import { NextResponse } from 'next/server';
import { grpcAsync } from '@/lib/grpc';

export async function POST() {
    try {
        const response = await grpcAsync('UpdateUserData', {});

        if (!response.general?.success) {
            return NextResponse.json(
                { error: response.general?.failure_details || 'Sync failed' },
                { status: 500 }
            );
        }

        return NextResponse.json({ success: true });
    } catch (error) {
        console.error('gRPC Error:', error);
        return NextResponse.json({ error: error.message }, { status: 500 });
    }
}
