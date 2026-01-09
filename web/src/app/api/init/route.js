import { NextResponse } from 'next/server';
import { grpcAsync } from '@/lib/grpc';

export async function GET() {
    try {
        const response = await grpcAsync('AuthenticateUser', {});

        if (!response.general?.success) {
            return NextResponse.json(
                { error: response.general?.failure_details || 'Authentication failed' },
                { status: 500 }
            );
        }

        return NextResponse.json({ userID: response.user_spotify_id });
    } catch (error) {
        console.error('gRPC Error:', error);
        return NextResponse.json({ error: error.message }, { status: 500 });
    }
}
