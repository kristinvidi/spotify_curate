import { NextResponse } from 'next/server';
import { grpcAsync } from '@/lib/grpc';

export async function POST(request) {
    const body = await request.json();
    const { userID, tag, artists } = body; // Frontend sends tag and list of artist names

    if (!userID || !tag || !artists) {
        return NextResponse.json({ error: 'Missing requirements' }, { status: 400 });
    }

    // Construct request matches GetGenreToArtistsMapping proto
    const payload = {
        user_spotify_id: userID,
        genre_to_artists_mappings: [
            {
                genre: tag,
                artist_names: artists,
            },
        ],
    };

    try {
        const response = await grpcAsync('CreateGenreToArtistsMappings', payload);

        if (!response.general?.success) {
            return NextResponse.json(
                { error: response.general?.failure_details || 'Failed to map artists' },
                { status: 500 }
            );
        }

        return NextResponse.json({
            success: true,
            failedMappings: response.failed_genre_to_artists_mappings
        });
    } catch (error) {
        console.error('gRPC Error:', error);
        return NextResponse.json({ error: error.message }, { status: 500 });
    }
}
